package ru.mephi.auth.service;

import java.io.ByteArrayInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.nio.charset.StandardCharsets;
import java.rmi.RemoteException;
import org.apache.commons.io.IOUtils;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;
import ru.mephi.auth.dto.CredDto;
import ru.mephi.auth.dto.PayloadDto;
import ru.mephi.auth.dto.PayloadWithTokenDto;
import ru.mephi.auth.dto.TokenDto;

@Service
public class CloudStoreApiHandler {
  private final RestTemplate restTemplate;
  private final String cloudStoreApiHost;

  public CloudStoreApiHandler(RestTemplate restTemplate, @Value("${cloud.store.host}") String cloudStoreApiHost) {
    this.restTemplate = restTemplate;
    this.cloudStoreApiHost = cloudStoreApiHost;
  }

  public TokenDto registerUser(String login, String pass) {
    var cred = new CredDto(login, pass);

    String url = cloudStoreApiHost + "/create_user";

    return restTemplate.postForObject(url, cred, TokenDto.class);
  }

  public TokenDto authUser(String login, String pass) {
    var cred = new CredDto(login, pass);

    String url = cloudStoreApiHost + "/auth_user";

    return restTemplate.postForObject(url, cred, TokenDto.class);
  }

  public InputStream loadKeyStore(String token) {
    String url = cloudStoreApiHost + "/get_payload";

    PayloadDto payload = restTemplate.postForObject(url, new TokenDto(token), PayloadDto.class);
    byte[] bytes = payload.payload().getBytes(StandardCharsets.UTF_8);

    return new ByteArrayInputStream(bytes);
  }

  public void uploadKeyStore(String token, InputStream keyStoreData) throws IOException {
    String url = cloudStoreApiHost + "/put_payload";

    PayloadWithTokenDto payloadWithTokenDto = new PayloadWithTokenDto(
        IOUtils.toString(keyStoreData, StandardCharsets.UTF_8),
        token
    );
    ResponseEntity<Void> response = restTemplate.postForEntity(url, payloadWithTokenDto, Void.class);

    if (!response.getStatusCode().is2xxSuccessful()) {
      throw new RemoteException();
    }
  }
}
