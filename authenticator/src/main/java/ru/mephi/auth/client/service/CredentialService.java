package ru.mephi.auth.client.service;

import java.security.KeyPairGenerator;
import java.security.KeyStore;
import org.springframework.stereotype.Service;
import ru.mephi.auth.client.exception.CredentialsException;
import ru.mephi.auth.client.model.Credential;

@Service
public class CredentialService {
  private final KeyPairGenerator keyPairGenerator;
  private final KeyStore keyStore;

  public CredentialService(KeyPairGenerator keyPairGenerator, KeyStore keyStore) {
    this.keyPairGenerator = keyPairGenerator;
    this.keyStore = keyStore;
  }

  public Credential getCredential(String login) throws CredentialsException {
    return null;
  }

  public Credential createCredential(String login) {
    return null;
  }
}
