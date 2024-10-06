
import React, { useState } from 'react';
import styles from '../styles/InputBox.module.scss';

const InputBox = ({ handleInput }) => {
  const [inputValue, setInputValue] = useState('');

  const handleChange = (e) => {
    setInputValue(e.target.value);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    handleInput(inputValue);
    setInputValue('');
  };

  return (
    <form onSubmit={handleSubmit} className={styles.inputContainer}>
      <input
        type="text"
        value={inputValue}
        onChange={handleChange}
        className={styles.inputField}
        placeholder="Ask Valeera..."
      />
      <button type="submit" className={styles.submitButton}>Send</button>
    </form>
  );
};

export default InputBox;
