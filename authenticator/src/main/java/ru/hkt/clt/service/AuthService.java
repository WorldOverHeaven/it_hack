package ru.hkt.clt.service;

import java.io.ByteArrayInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.math.BigInteger;
import java.security.KeyPair;
import java.security.KeyPairGenerator;
import java.security.KeyStoreException;
import java.security.NoSuchAlgorithmException;
import java.security.PrivateKey;
import java.security.SecureRandom;
import java.security.Signature;
import java.security.cert.CertificateException;
import java.security.cert.X509Certificate;
import java.util.Base64;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;
import javax.security.auth.x500.X500Principal;
import org.bouncycastle.cert.X509CertificateHolder;
import org.bouncycastle.cert.jcajce.JcaX509CertificateConverter;
import org.bouncycastle.cert.jcajce.JcaX509v3CertificateBuilder;
import org.bouncycastle.jce.provider.BouncyCastleProvider;
import org.bouncycastle.operator.ContentSigner;
import org.bouncycastle.operator.jcajce.JcaContentSignerBuilder;
import org.springframework.stereotype.Service;
import ru.hkt.clt.dto.ChallengeDto;
import ru.hkt.clt.dto.MessageDto;

@Service
public class AuthService {
  private volatile String cloudToken;
  private final KeyStoreService keyStoreService;
  private final CloudStoreApiHandler cloudStoreApiHandler;
  private final WebAuthServiceApiHandler webAuthServiceApiHandler;
  private final KeyPairGenerator keyPairGenerator;
  private final Map<String, String> tokens = new HashMap<>();

  public AuthService(KeyStoreService keyStoreService, CloudStoreApiHandler cloudStoreApiHandler, WebAuthServiceApiHandler webAuthServiceApiHandler,
      KeyPairGenerator keyPairGenerator
  ) {
    this.keyStoreService = keyStoreService;
    this.cloudStoreApiHandler = cloudStoreApiHandler;
    this.webAuthServiceApiHandler = webAuthServiceApiHandler;
    this.keyPairGenerator = keyPairGenerator;
  }

  public void joinCloud(String cloudLogin, String cloudPass) {
    try {
      this.cloudToken = cloudStoreApiHandler.authUser(cloudLogin, cloudPass).token();
    } catch (Exception e) {
      this.cloudToken = cloudStoreApiHandler.registerUser(cloudLogin, cloudPass).token();
    }
  }

  public void registerNewUser(String login) throws Exception {
    if (cloudToken == null) {
      throw new RuntimeException();
    }
    KeyPair keyPair = keyPairGenerator.generateKeyPair();
    String token = webAuthServiceApiHandler.createUser(login, keyPair.getPublic()).token();
    tokens.put(login, token);
    keyStoreService.putKeys(keyPair, generateCertificate(keyPair), login);
    synchronizeKeyStoreWithCloud();
  }

  public void authUser(String login) throws Exception {
    if (cloudToken == null) {
      throw new RuntimeException();
    }
    synchronizeKeyStoreWithCloud();
    KeyPair keyPair = keyStoreService.getKeyPair(login);
    ChallengeDto challenge = webAuthServiceApiHandler.getChallenge(login, keyPair.getPublic());
    String sign = signChallenge(challenge.challenge(), keyPair.getPrivate());
    String token = webAuthServiceApiHandler.solveChallenge(challenge.challenge_id(), sign).token();
    tokens.put(login, token);
  }

  public MessageDto verify(String login) {
    if (cloudToken == null) {
      throw new RuntimeException();
    }
    return webAuthServiceApiHandler.verify(tokens.get(login));
  }

  private void synchronizeKeyStoreWithCloud() throws IOException, CertificateException, KeyStoreException, NoSuchAlgorithmException {
    byte[] cloudData = cloudStoreApiHandler.loadKeyStore(cloudToken);
    if (cloudData.length > 0) {
      try (InputStream is =new ByteArrayInputStream(cloudData)) {
        keyStoreService.addKeysFromOtherKeyStore(is);
      } catch (Exception e) {
        throw new RuntimeException(e);
      }
    }
    cloudStoreApiHandler.uploadKeyStore(cloudToken, keyStoreService.getKeyStore());
  }

  private X509Certificate generateCertificate(KeyPair keyPair) throws Exception {
    Date from = new Date();
    Date to = new Date(from.getTime() + 365 * 86400000L); // 1 year validity
    BigInteger sn = new BigInteger(64, new SecureRandom());
    ContentSigner signer = new JcaContentSignerBuilder("SHA256WithRSAEncryption").build(keyPair.getPrivate());

    JcaX509v3CertificateBuilder builder = new JcaX509v3CertificateBuilder(
        new X500Principal("CN=example.com"), sn, from, to, new X500Principal("CN=example.com"), keyPair.getPublic());

    X509CertificateHolder certHolder = builder.build(signer);
    return new JcaX509CertificateConverter().setProvider(new BouncyCastleProvider()).getCertificate(certHolder);
  }

  private static String signChallenge(String challenge, PrivateKey privateKey) throws Exception {
    byte[] challengeBytes = Base64.getDecoder().decode(challenge);
    Signature signature = Signature.getInstance("SHA256withRSA");
    signature.initSign(privateKey);
    signature.update(challengeBytes);
    return Base64.getEncoder().encodeToString(signature.sign());
  }
}
