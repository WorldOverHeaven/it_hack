package ru.mephi.auth.client.entity;

import org.hibernate.Hibernate;

import javax.persistence.Column;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.Entity;
import javax.persistence.Table;
import java.util.Objects;

@Entity
@Table(name = "task")
public class Task {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "task_id")
    private Long id;

    @Column(
            name = "title",
            nullable = false)
    private String title;

    @Column(
            name = "completion_status",
            nullable = false)
    private Boolean completionStatus;

    public Task() {
    }

    public Task(String title, Boolean completionStatus) {
        this.title = title;
        this.completionStatus = completionStatus;
    }

    public Task(Long id, String title, Boolean completionStatus) {
        this.id = id;
        this.title = title;
        this.completionStatus = completionStatus;
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public Boolean getCompletionStatus() {
        return completionStatus;
    }

    public void setCompletionStatus(Boolean completionStatus) {
        this.completionStatus = completionStatus;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || Hibernate.getClass(this) != Hibernate.getClass(o)) return false;
        Task that = (Task) o;
        return getId() != null && Objects.equals(getId(), that.getId());
    }

    @Override
    public int hashCode() {
        return getClass().hashCode();
    }
}
