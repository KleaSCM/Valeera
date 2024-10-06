import React, { useState } from 'react';
import styles from '../styles/settings.module.scss'; 

const SettingsPage = () => {
  const [question, setQuestion] = useState('');
  const [response, setResponse] = useState('');
  const [message, setMessage] = useState('');
  const [success, setSuccess] = useState(false); // New state to track success

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    const data = {
      question,
      response,
    };

    // Updated the fetch URL to point to the Go backend
    const res = await fetch('http://localhost:8080/api/updateSnippets', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });

    if (res.ok) {
      setMessage('Successfully added!');
      setSuccess(true); // Mark as success
      setQuestion('');
      setResponse('');
    } else {
      setMessage('Error adding the response.');
      setSuccess(false); // Mark as failure
    }
  };

  return (
    <div className={styles.settingsContainer}>
      <h1 className={styles.settingsTitle}>Settings</h1>
      <form className={styles.settingsForm} onSubmit={handleSubmit}>
        <div className={styles.formField}>
          <label htmlFor="question">Question:</label>
          <input
            type="text"
            id="question"
            value={question}
            onChange={(e) => setQuestion(e.target.value)}
          />
        </div>
        <div className={styles.formField}>
          <label htmlFor="response">Response:</label>
          <textarea
            id="response"
            value={response}
            onChange={(e) => setResponse(e.target.value)}
          />
        </div>
        <button className={styles.submitButton} type="submit">Add</button>
      </form>
      {message && (
        <p className={`${styles.message} ${success ? styles.success : styles.error}`}>
          {message}
        </p>
      )}
    </div>
  );
};

export default SettingsPage;
