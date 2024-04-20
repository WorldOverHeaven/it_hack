package ru.mephi.auth.service;

import java.io.InputStream;
import org.springframework.stereotype.Service;
import ru.mephi.auth.dto.TokenDto;

@Service
public class CloudStoreApiHandler {
  public TokenDto registerUser(String login, String pass) {
    return null;
  }

  public TokenDto authUser(String login, String pass) {
    return null;
  }

  public InputStream loadKeyStore(String token) {
    return null;
  }

  public void uploadKeyStore(String token, InputStream stream) {

  }
}
