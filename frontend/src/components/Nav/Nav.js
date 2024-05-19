import React from 'react';
import './Nav.css';
import { useOptionContext } from '../Context/Context';

function Nav() {
  const { options, setOptions } = useOptionContext();

  const handleOptionChange = (option) => {
    setOptions({ ...options, selectedOption: option,key:'',value:'', expiryTime:'',showCard:false });
  };

  return (
    <div className='navbar-container'>
      <p className={options.selectedOption === 'get' ? 'active' : ''} onClick={() => {handleOptionChange('get')}}>Retrieve</p>
      <p className={options.selectedOption === 'post' ? 'active' : ''} onClick={() => handleOptionChange('post')}>Save</p>
      <p className={options.selectedOption === 'delete' ? 'active' : ''} onClick={() => handleOptionChange('delete')}>Remove</p>
    </div>
  );
}

export default Nav;
