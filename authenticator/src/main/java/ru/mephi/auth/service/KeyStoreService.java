package ru.mephi.auth.service;

import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.security.Key;
import java.security.KeyPair;
import java.security.KeyStore;
import java.security.KeyStoreException;
import java.security.NoSuchAlgorithmException;
import java.security.PrivateKey;
import java.security.PublicKey;
import java.security.UnrecoverableKeyException;
import java.security.cert.Certificate;
import java.security.cert.CertificateException;
import java.security.cert.X509Certificate;
import java.util.Base64;
import java.util.Enumeration;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

@Service
public class KeyStoreService {
  private final String keyStorePassword;
  private final String keyStorePath;
  private final KeyStore keyStore;
  private static final Logger LOG = LoggerFactory.getLogger(KeyStoreService.class);

  public KeyStoreService(@Value("${keystore.password}") String keyStorePassword, @Value("${keystore.path}") String keyStorePath, KeyStore keyStore) {
    this.keyStorePassword = keyStorePassword;
    this.keyStorePath = keyStorePath;
    this.keyStore = keyStore;
  }

  public KeyPair getKeyPair(String login) throws UnrecoverableKeyException, KeyStoreException, NoSuchAlgorithmException {
    return new KeyPair(
        getPublicKey(login),
        getPrivateKey(login)
    );
  }

  public PrivateKey getPrivateKey(String login) throws UnrecoverableKeyException, KeyStoreException, NoSuchAlgorithmException {
    if (!keyStore.containsAlias(login)) {
      throw new KeyStoreException("No key found for alias: " + login);
    }

    PrivateKey privateKey = (PrivateKey) keyStore.getKey(login, keyStorePassword.toCharArray());
    if (privateKey == null) {
      throw new KeyStoreException("No private key found for alias: " + login);
    }

    return privateKey;
  }

  public PublicKey getPublicKey(String login) throws KeyStoreException {
    Certificate cert = keyStore.getCertificate(login);
    if (cert == null) {
      throw new KeyStoreException("No certificate found for alias: " + login);
    }

    return cert.getPublicKey();
  }

  public String getPublicKeyString(String login) throws KeyStoreException {
    Certificate cert = keyStore.getCertificate(login);
    if (cert == null) {
      throw new KeyStoreException("No certificate found for alias: " + login);
    }

    return Base64.getEncoder().encodeToString(getPublicKey(login).getEncoded());
  }

  public void putKeys(
      KeyPair keyPair,
      X509Certificate certificate,
      String login
  ) throws UnrecoverableKeyException, KeyStoreException, NoSuchAlgorithmException, CertificateException, IOException {
    X509Certificate[] certChain = new X509Certificate[]{certificate};
    keyStore.setKeyEntry(login, keyPair.getPrivate(), keyStorePassword.toCharArray(), certChain);

    try (FileOutputStream fos = new FileOutputStream(keyStorePath)) {
      keyStore.store(fos, keyStorePassword.toCharArray());
    }
  }

  public InputStream getKeyStore() throws FileNotFoundException {
    FileInputStream fis;
    try {
      fis = new FileInputStream(keyStorePath);
    } catch (FileNotFoundException e) {
      throw new FileNotFoundException("KeyStore file not found at " + keyStorePath);
    }
    return fis;
  }

  public void addKeysFromOtherKeyStore(InputStream keyStoreData) throws KeyStoreException, CertificateException, IOException, NoSuchAlgorithmException, UnrecoverableKeyException {
    KeyStore externalKeyStore = KeyStore.getInstance(keyStore.getType());

    externalKeyStore.load(keyStoreData, keyStorePassword.toCharArray());

    for (Enumeration<String> aliases = externalKeyStore.aliases(); aliases.hasMoreElements(); ) {
      String alias = aliases.nextElement();
      if (!keyStore.containsAlias(alias)) {
        Key key = externalKeyStore.getKey(alias, keyStorePassword.toCharArray());
        Certificate[] certChain = externalKeyStore.getCertificateChain(alias);

        keyStore.setKeyEntry(alias, key, keyStorePassword.toCharArray(), certChain);
      } else {
        LOG.info("Alias {} already exists in the local key store.", alias);
      }
    }

    try (FileOutputStream fos = new FileOutputStream(keyStorePath)) {
      keyStore.store(fos, keyStorePassword.toCharArray());
    }
  }
}
