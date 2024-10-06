import React from 'react';
import styles from '../styles/OutputBox.module.scss';

type OutputBoxProps = {
  response: string; 
};

const OutputBox: React.FC<OutputBoxProps> = ({ response }) => {
  return (
    <div className={styles.outputContainer}>
      <p className={styles.responseText}>{response}</p>
    </div>
  );
};

export default OutputBox;
