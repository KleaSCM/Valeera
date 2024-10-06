import React, { useState } from 'react';
import Image from 'next/image';
import InputBox from './InputBox';
import OutputBox from './OutputBox';
import { sendMessage } from '../utils/api/message'; // Import the sendMessage function
import styles from '../styles/ValeeraCard.module.scss';

const ValeeraCard = () => {
    const [responseMessage, setResponseMessage] = useState('Hello! I am Valeera, your assistant.');
    const [codeSnippet, setCodeSnippet] = useState<string | undefined>(undefined); // State for code snippet

    const handleSendMessage = async (message: string) => {
        try {
            const data = await sendMessage(message); // Call the sendMessage function

            // Update response and code from backend
            setResponseMessage(data.reply);
            setCodeSnippet(data.code); // Assume the backend returns code if applicable
        } catch (error) {
            console.error('Error sending message:', error);
            setResponseMessage('Sorry, something went wrong.');
            setCodeSnippet(undefined); // Reset code snippet on error
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
            <OutputBox response={responseMessage} code={codeSnippet} /> {/* Pass code snippet */}
        </div>
    );
};

export default ValeeraCard;
