
import React, { useState } from 'react';
import Image from 'next/image';
import InputBox from './InputBox';
import OutputBox from './OutputBox';
import styles from '../styles/ValeeraCard.module.scss';

const ValeeraCard = () => {
  const [response, setResponse] = useState('Hello, I’m Valeera! What can I help you with?');

  const handleInput = (input) => {
    // echo input as response change logic later.
    setResponse(`Valeera: ${input}`);
  };

  return (
    <div className={styles.card}>
      <h1 className={styles.title}>Valeera</h1>
      <Image
        src="/valeera.jpeg"
        alt="Valeera"
        width={400}
        height={600}
        className={styles.valeeraImage}
      />
      <InputBox handleInput={handleInput} />
      <OutputBox response={response} />
    </div>
  );
};

export default ValeeraCard;
