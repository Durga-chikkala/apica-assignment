import React from 'react'
import Logo from '../../Images/logo.png'
import './Loader.css'
function Loader() {
  return (
    <div className='loader'>
        <img src={Logo} alt="Loading..."/>
    </div>
  )
}

export default Loader