import React, { useState, useEffect } from 'react';
import './GetKeys.css'

function Home() {
  const [cacheData, setCacheData] = useState({});

  useEffect(() => {
    const ws = new WebSocket('ws://localhost:8080/keys');

    ws.onopen = () => {
      console.log('WebSocket connected');
    };

    ws.onmessage = (event) => {
      const data = JSON.parse(event.data);
      setCacheData(data);
    };

    ws.onerror = (error) => {
      console.error('WebSocket error:', error);
    };

    ws.onclose = () => {
      console.log('WebSocket closed');
    };

    return () => {
      ws.close();
    };
  }, []);

  return (
    <div className="home">
    <h1>Cache Data</h1>
    <div className="cache-container">
      {Object.keys(cacheData).map((key) => (
        <div className="cache-card" key={key}>
          <p><strong>Key:</strong> {cacheData[key].key}</p>
          <p><strong>Value:</strong> {cacheData[key].value}</p>
          <p><strong>Expiration:</strong> {cacheData[key].expiration}</p>
        </div>
      ))}
    </div>
  </div>
  );
}

export default Home;
