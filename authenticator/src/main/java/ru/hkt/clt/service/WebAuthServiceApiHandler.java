package ru.hkt.clt.service;

import java.security.PublicKey;
import java.util.Base64;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;
import ru.hkt.clt.dto.ChallengeDto;
import ru.hkt.clt.dto.LoginWithPublicKeyDto;
import ru.hkt.clt.dto.MessageDto;
import ru.hkt.clt.dto.SolvedChallengeDto;
import ru.hkt.clt.dto.TokenDto;

@Service
public class WebAuthServiceApiHandler {
  private final RestTemplate restTemplate;
  private final String webAuthServiceApiHost;

  public WebAuthServiceApiHandler(RestTemplate restTemplate, @Value("${web.auth.service.host}") String webAuthServiceApiHost) {
    this.restTemplate = restTemplate;
    this.webAuthServiceApiHost = webAuthServiceApiHost;
  }

  public TokenDto createUser(String login, PublicKey publicKey) {
    LoginWithPublicKeyDto request = new LoginWithPublicKeyDto(login, toStr(publicKey));
    String url = webAuthServiceApiHost + "/create_user";

    return restTemplate.postForObject(url, request, TokenDto.class);
  }

  public ChallengeDto getChallenge(String login, PublicKey publicKey) {
    LoginWithPublicKeyDto request = new LoginWithPublicKeyDto(login, toStr(publicKey));
    String url = webAuthServiceApiHost + "/get_challenge";

    return restTemplate.postForObject(url, request, ChallengeDto.class);
  }

  public TokenDto solveChallenge(String challengeId, String solvedChallenge) {
    SolvedChallengeDto request = new SolvedChallengeDto(challengeId, solvedChallenge);
    String url = webAuthServiceApiHost + "/solve_challenge";

    return restTemplate.postForObject(url, request, TokenDto.class);
  }

  public MessageDto verify(String token) {
    TokenDto request = new TokenDto(token);
    String url = webAuthServiceApiHost + "/verify";

    return restTemplate.postForObject(url, request, MessageDto.class);
  }

  private String toStr(PublicKey publicKey) {
    return Base64.getEncoder().encodeToString(publicKey.getEncoded());
  }
}
