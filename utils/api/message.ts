export const sendMessage = async (content: string) => {
    const response = await fetch("http://localhost:8080/api/message", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ content }),
    });

    if (!response.ok) {
        const errorMessage = await response.text(); // Get the response text for debugging
        throw new Error(`Failed to fetch response from the server: ${errorMessage}`);
    }

    return response.json();
};
