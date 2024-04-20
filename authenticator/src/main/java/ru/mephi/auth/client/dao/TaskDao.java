package ru.mephi.auth.client.dao;

import jakarta.inject.Inject;
import org.hibernate.Session;
import org.hibernate.SessionFactory;
import org.springframework.stereotype.Repository;
import ru.mephi.auth.client.entity.Task;

import javax.persistence.criteria.CriteriaBuilder;
import javax.persistence.criteria.CriteriaQuery;
import javax.persistence.criteria.Predicate;
import javax.persistence.criteria.Root;
import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

@Repository
public class TaskDao {

    private final SessionFactory sessionFactory;

    @Inject
    public TaskDao(SessionFactory sessionFactory){
        this.sessionFactory = sessionFactory;
    }

    public List<Task> find(Long id, String title, Boolean completionStatus){
        Session session = getSession();
        CriteriaBuilder builder = session.getCriteriaBuilder();
        CriteriaQuery<Task> criteriaQuery = builder.createQuery(Task.class);
        Root<Task> root = criteriaQuery.from(Task.class);

        List<Predicate> predicates = new ArrayList<>();
        if(id != null) {
            predicates.add(builder.equal(root.get("id"), id));
        }
        if(title != null) {
            predicates.add(builder.equal(root.get("title"), title));
        }
        if(completionStatus != null){
            predicates.add(builder.equal(root.get("completionStatus"), completionStatus));
        }
        criteriaQuery.select(root).where(predicates.toArray(new Predicate[]{}));
        return session.createQuery(criteriaQuery).getResultList();
    }

    public Optional<Task> getTask(Long id){
        return Optional.ofNullable(getSession().get(Task.class, id));
    }

    public void update(Task task){
        getSession().update(task);
    }

    public Task save(String title, Boolean completionStatus){
        Task task = new Task(title, completionStatus);
        getSession().persist(task);
        return task;
    }

    public void remove(Long id){
        getSession()
                .createQuery("DELETE Task t WHERE t.id = :id")
                .setParameter("id", id)
                .executeUpdate();
    }

    public void removeAll(){
        getSession()
                .createQuery("DELETE Task")
                .executeUpdate();
    }

    private Session getSession() {
        return sessionFactory.getCurrentSession();
    }
}
