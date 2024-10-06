import React, { useState } from 'react';
import Image from 'next/image';
import InputBox from './InputBox';
import OutputBox from './OutputBox';
import styles from '../styles/ValeeraCard.module.scss';

const ValeeraCard = () => {
  const [responseMessage, setResponseMessage] = useState('Hello! I am Valeera, your assistant.');

  const handleSendMessage = async (message: string) => {
    try {
      const res = await fetch('http://localhost:8080/api/message', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ content: message }),
      });

      const data = await res.json();
      setResponseMessage(data.reply); // Update response from backend
    } catch (error) {
      console.error('Error sending message:', error);
      setResponseMessage('Sorry, something went wrong.');
    }
  };

  return (
    <div className={styles.card}>
      <h1 className={styles.heading}>Valeera</h1>
      <Image
        src="/valeera.jpeg" 
        alt="Valeera"
        width={400}
        height={600}
        className={styles.image}
      />
      {/* Message Input */}
      <InputBox handleInput={handleSendMessage} />
      {/* Message Output */}
      <OutputBox response={responseMessage} />
    </div>
  );
};

export default ValeeraCard;
