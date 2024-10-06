

import type { NextApiRequest, NextApiResponse } from 'next';

// (Replace with actual OpenAI integration)
const mockOpenAiResponse = async (question: string) => {
    return {
        answer: `This is a mock response to your question: ${question}`
    };
};

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
    const { question } = req.body;

    if (!question) {
        return res.status(400).json({ error: 'No question provided' });
    }

    // Replace  with actual OpenAI API call
    const response = await mockOpenAiResponse(question);
    res.status(200).json(response);
}
