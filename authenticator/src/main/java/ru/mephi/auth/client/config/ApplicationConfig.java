package ru.mephi.auth.client.config;

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
import org.glassfish.jersey.servlet.ServletProperties;
import org.hibernate.SessionFactory;
import org.hibernate.cfg.Environment;
import org.hibernate.dialect.PostgreSQL10Dialect;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.jdbc.datasource.DriverManagerDataSource;
import org.springframework.orm.hibernate5.HibernateTransactionManager;
import org.springframework.orm.hibernate5.LocalSessionFactoryBean;
import ru.mephi.auth.client.entity.Task;
import ru.mephi.auth.client.resource.TaskResource;

import javax.sql.DataSource;
import java.util.Properties;

@Configuration
public class ApplicationConfig {
    @Value("${task.db.url}")
    private String taskDbUrl;
    @Value("${task.db.user}")
    private String taskDbUser;
    @Value("${task.db.password}")
    private String taskDbPassword;
    @Value("${keystore.path}")
    private String keyStorePath;
    @Value("${keystore.password}")
    private String keyStorePassword;

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
    public DataSource dataSource() {
        return new DriverManagerDataSource(taskDbUrl, taskDbUser, taskDbPassword);
    }

    @Bean
    LocalSessionFactoryBean localSessionFactoryBean(DataSource dataSource) {
        LocalSessionFactoryBean localSessionFactoryBean = new LocalSessionFactoryBean();
        localSessionFactoryBean.setDataSource(dataSource);

        Properties properties = new Properties();
        properties.put(Environment.DIALECT, PostgreSQL10Dialect.class.getName());
        properties.put(Environment.HBM2DDL_AUTO, "update");
        localSessionFactoryBean.setHibernateProperties(properties);
        localSessionFactoryBean.setPackagesToScan("dao");
        localSessionFactoryBean.setAnnotatedClasses(Task.class);
        return localSessionFactoryBean;
    }

    @Bean
    public SessionFactory sessionFactory(LocalSessionFactoryBean localSessionFactoryBean) {
        return localSessionFactoryBean.getObject();
    }

    @Bean
    public HibernateTransactionManager platformTransactionManager(SessionFactory sessionFactory) {
        return new HibernateTransactionManager(sessionFactory);
    }

    @Bean
    public ResourceConfig jerseyConfig() {
        ResourceConfig config = new ResourceConfig();
        config.register(TaskResource.class);
        config.property(ServletProperties.FILTER_FORWARD_ON_404, true);
        return config;
    }
}
