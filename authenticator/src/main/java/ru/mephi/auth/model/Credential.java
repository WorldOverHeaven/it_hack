package ru.mephi.auth.model;

public record Credential(
    String openKey,
    String privateKey
) {
}
