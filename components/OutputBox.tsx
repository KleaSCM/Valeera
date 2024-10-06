
import React from 'react';
import styles from '../styles/OutputBox.module.scss';

const OutputBox = ({ response }) => {
  return (
    <div className={styles.outputContainer}>
      <p className={styles.responseText}>{response}</p>
    </div>
  );
};

export default OutputBox;
