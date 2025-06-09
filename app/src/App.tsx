import { useState } from 'react';
import { createConnectTransport } from '@connectrpc/connect-web';
import { createPromiseClient } from '@connectrpc/connect';
import { GreetService } from '../gen/greet/v1/greet_connectweb';

const transport = createConnectTransport({
  baseUrl: 'http://localhost:8080',
});

function App() {
  const [response, setResponse] = useState('');
  const [error, setError] = useState('');
  const client = createPromiseClient(GreetService, transport);

  const handleGreet = async () => {
    try {
      setError('');
      const result = await client.greet({ name: 'Guest' });
      setResponse(JSON.stringify(result, null, 2));
    } catch (err) {
      setError(err instanceof Error ? err.message : String(err));
    }
  };

  const handleGreetStream = async () => {
    try {
      setError('');
      setResponse('');
      let messages = '';
      
      for await (const result of client.greetStream({ name: 'Guest' })) {
        messages += JSON.stringify(result, null, 2) + '\n\n';
        setResponse(messages);
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : String(err));
    }
  };

  return (
    <div style={{ padding: '20px', maxWidth: '800px', margin: '0 auto' }}>
      <h1>Connect-Web Greeter</h1>
      <div style={{ marginBottom: '20px' }}>
        <button onClick={handleGreet} style={{ marginRight: '10px' }}>Greet</button>
        <button onClick={handleGreetStream}>Start Stream</button>
      </div>
      {error && (
        <div style={{ color: 'red', marginBottom: '10px' }}>{error}</div>
      )}
      {response && (
        <pre style={{ 
          background: '#f5f5f5', 
          padding: '15px', 
          borderRadius: '4px',
          whiteSpace: 'pre-wrap'
        }}>
          {response}
        </pre>
      )}
    </div>
  );
}

export default App;
