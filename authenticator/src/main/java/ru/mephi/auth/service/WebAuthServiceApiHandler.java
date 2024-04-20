package ru.mephi.auth.service;

import java.security.PublicKey;
import org.springframework.stereotype.Service;
import ru.mephi.auth.dto.ChallengeDto;
import ru.mephi.auth.dto.MessageDto;
import ru.mephi.auth.dto.TokenDto;

@Service
public class WebAuthServiceApiHandler {
  public TokenDto createUser(String login, PublicKey publicKey) {
    return null;
  }

  public ChallengeDto getChallenge(String login, PublicKey publicKey) {
    return null;
  }

  public TokenDto solveChallenge(String challengeId, String solvedChallenge) {
    return null;
  }

  public MessageDto verify(String token) {
    return null;
  }
}
