import React, { createContext, useContext, useState } from 'react';

const OptionContext = createContext();

export const useOptionContext = () => useContext(OptionContext);

export const OptionProvider = ({ children }) => {
  const [options, setOptions] = useState({
    selectedOption: 'get',
    key: '',
    value: '',
    expiryTime: '',
    showCard: false
  });

  return (
    <OptionContext.Provider value={{ options, setOptions }}>
      {children}
    </OptionContext.Provider>
  );
};
