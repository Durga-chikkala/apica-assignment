import './App.css';
import Home  from './components/GetALLKeys/GetKeys';
import Header from './components/Header/Header';
import Nav from './components/Nav/Nav';


function App() {
  return (
    <div className="App">
        <Nav/>
        <Header/>
        <Home/>
    </div>
  );
}

export default App;
