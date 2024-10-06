import React, { useState } from 'react';
import styles from '../styles/InputBox.module.scss';

type InputBoxProps = {
  handleInput: (message: string) => void; 
};

const InputBox: React.FC<InputBoxProps> = ({ handleInput }) => {
  const [inputValue, setInputValue] = useState('');

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(e.target.value);
  };

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (inputValue.trim()) {
      handleInput(inputValue);  // Send input to parent handler
      setInputValue('');        // Clear the input after submission
    }
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
