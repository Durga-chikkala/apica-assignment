import React from 'react';
import useAxios from '../useHook/useHook';
import './Header.css';
import Loader from '../Loader/Loader';
import { useOptionContext } from '../Context/Context';

const Header = () => {
  const { options, setOptions } = useOptionContext();
  const { selectedOption } = options;

  const { response, error, isLoading, fetchData,setResponse,setError,setIsLoading  } = useAxios();

  const handleGetSubmit = () => {
    setOptions({ ...options,showCard:true });
    setResponse(null)
    setError(null)
    setIsLoading(false)

    if (options.key === "") {
      setError( "Key cannot be empty");
      return;
    }

    fetchData('GET', `http://localhost:8080/cache/${options.key}`);
  };

  const handleDeleteSubmit = () => {
    setOptions({ ...options,showCard:true });
    setResponse(null)
    setError(null)
    setIsLoading(false)

    if (options.key === "") {
      setError( "Key cannot be empty");
      return;
    }

    fetchData('DELETE', `http://localhost:8080/cacheKey/${options.key}`);
  };

  const handleSetSubmit = () => {
    setOptions({ ...options,showCard:true });
    setResponse(null)
    setError(null)
    setIsLoading(false)

    if (options.key === "") {
      setError( "Key cannot be empty");
      return;
    }

    if (options.value === "") {
      setError("Value cannot be empty");
      return
    }

    fetchData('PUT', 'http://localhost:8080/cache', { "key": options.key, 
    "value": options.value, "expiration": parseInt(options.expiration, 10) });
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setOptions({ ...options, [name]: value });
  };

  return (
    <div className="header-container">
      <h1>CacheSmart</h1>
      <h3> Lightning-Fast In-Memory Solutions</h3>
      <div className="header-form">
        {selectedOption === 'get' || selectedOption === "delete" ? (
          <div>
            <input
              type="text"
              value={options.key}
              onChange={handleChange}
              name="key"
              placeholder="Enter Key"
              required
            />
            <button onClick={selectedOption === "get" ? handleGetSubmit : handleDeleteSubmit}>{selectedOption === "get" ? "Get Data" : "Delete Data"}</button>
          </div>
        ) : (
          <div>
            <input
              type="text"
              value={options.key}
              onChange={handleChange}
              name="key"
              placeholder="Key"
              required
            />
            <input
              type="text"
              value={options.value}
              onChange={handleChange}
              name="value"
              placeholder="Value"
              required
            />

            <input
              type="text"
              value={options.expiration}
              onChange={handleChange}
              name="expiration"
              placeholder="Expiry"
              required
            />
            <button onClick={handleSetSubmit}>Set Data</button>
          </div>
        )}
      </div>
      <div className={`header-response ${options.showCard ? 'show' : ''}`}>
        {isLoading && <p><Loader /></p>}
        {error && <p>Error: {error}</p>}
        {response && <pre> {JSON.stringify(response.data)}</pre>}
      </div>
    </div>
  );
};

export default Header;
