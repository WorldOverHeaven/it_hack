package ru.mephi.auth.client.exception;

import javassist.NotFoundException;

public class CredentialsException extends RuntimeException {
  public CredentialsException(String message) {
    super(message);
  }
}
