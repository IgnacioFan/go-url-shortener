import './App.css'
import Header from './components/header';
import Form from "./components/form";
import Footer from "./components/footer";
import List from './components/list';

function App() {
  return (
    <>
      <div className='body'>
        <Header/>
        <Form/>
        <List/>
        <Footer/>
      </div>
    </>
  )
}

export default App
