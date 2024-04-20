package ru.mephi.auth.exception;

import javassist.NotFoundException;

public class CredentialsException extends RuntimeException {
  public CredentialsException(String message) {
    super(message);
  }
}
