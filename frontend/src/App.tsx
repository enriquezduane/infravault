import { useState, useEffect } from "react";

function App() {
  const [message, setMessage] = useState<string>("Loading...");

  useEffect(() => {
    const fetchHelloWorld = async () => {
      try {
        const response = await fetch("/api");
        const data = await response.text();
        setMessage(data);
      } catch (error) {
        console.error("Error fetching data:", error);
        setMessage("Error connecting to backend");
      }
    };

    fetchHelloWorld();
  }, []);

  return (
    <>
      <h1>{message}</h1>
      <p>this is from the client!</p>
    </>
  );
}

export default App;
