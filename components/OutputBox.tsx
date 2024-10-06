import React, { useState } from 'react';
import styles from '../styles/OutputBox.module.scss';

type OutputBoxProps = {
  response: string;
  code?: string; 
};

const OutputBox: React.FC<OutputBoxProps> = ({ response, code }) => {
  const [isExpanded, setIsExpanded] = useState(false);

  const handleCopy = () => {
    if (code) {
      navigator.clipboard.writeText(code).then(() => {
        alert('Code copied to clipboard!');
      }).catch(err => {
        console.error('Failed to copy: ', err);
      });
    }
  };

  const toggleExpand = () => {
    setIsExpanded((prev) => !prev);
  };

  return (
    <div className={styles.outputContainer}>
      <p className={styles.responseText}>{response}</p>
      {code && (
        <div className={`${styles.codeContainer} ${isExpanded ? styles.expanded : ''}`}>
          <pre className={styles.codeSnippet}>
            <code>{code}</code>
          </pre>
          <button className={styles.toggleButton} onClick={toggleExpand}>
            {isExpanded ? 'Show Less' : 'Show More'}
          </button>
          <button className={styles.copyButton} onClick={handleCopy}>
            Copy Code
          </button>
        </div>
      )}
    </div>
  );
};

export default OutputBox;

