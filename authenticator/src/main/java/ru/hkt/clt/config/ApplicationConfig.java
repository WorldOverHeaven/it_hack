package ru.hkt.clt.config;

import java.io.File;
import java.io.FileInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.security.KeyPairGenerator;
import java.security.KeyStore;
import java.security.KeyStoreException;
import java.security.NoSuchAlgorithmException;
import java.security.cert.CertificateException;
import org.glassfish.jersey.server.ResourceConfig;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.client.RestTemplate;
import ru.hkt.clt.resource.AuthResource;

@Configuration
public class ApplicationConfig {
    @Value("${keystore.path}")
    private String keyStorePath;
    @Value("${keystore.password}")
    private String keyStorePassword;

    @Bean
    public RestTemplate restTemplate() {
        return new RestTemplate();
    }

    @Bean
    public KeyPairGenerator keyPairGenerator() throws NoSuchAlgorithmException {
        return KeyPairGenerator.getInstance("RSA");
    }

    @Bean
    public KeyStore getKeyStore() throws KeyStoreException, CertificateException, IOException, NoSuchAlgorithmException {
        KeyStore keyStore = KeyStore.getInstance(KeyStore.getDefaultType());
        File keystoreFile = new File(keyStorePath);
        if (keystoreFile.exists()) {
            try (InputStream is = new FileInputStream(keystoreFile)) {
                keyStore.load(is, keyStorePassword.toCharArray());
            }
        } else {
            keyStore.load(null, keyStorePassword.toCharArray());
        }
        return keyStore;
    }



    @Bean
    public ResourceConfig jerseyConfig() {
        ResourceConfig config = new ResourceConfig();
        config.register(AuthResource.class);
        return config;
    }
}
