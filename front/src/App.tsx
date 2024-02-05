import "./App.css";
import { useState } from "react";
import Header from "./components/header";
import Form from "./components/form";
import Footer from "./components/footer";
import List from "./components/list";
import { Url } from "./interfaces/url";
import { formData } from "./interfaces/form";

const urlRegex = /^(https?|ftp):\/\/(-\.)?([^\s/?\.#-]+\.?)+([^\s]*)$/i;

function App() {
  const [form, setForm] = useState<formData>({
    longUrl: ""
  });

  const [showWarning, setShowWarning] = useState<string>("");

  const [urls, setUrls] = useState<Url[]>(
    [
      {
        longUrl: "https://www.udemy.com/",
        shortUrl: "abc",
      },
      {
        longUrl: "https://www.kkday.com/zh-tw",
        shortUrl: "edc",
      }
    ]
  )

  const urlInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target;
    setShowWarning("")
    setForm({ ...form, [name]: value });
  }

  const genShortUrl = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    if (form.longUrl === "" || !validateUrl(form.longUrl)) {
      setShowWarning("Please enter a valid URL link!")
      return
    }
    if (urls.find(pair => pair.longUrl === form.longUrl)) {
      setShowWarning("The given URL already exists!")
      return
    }
    // try making a request and catch exception
    let pair: Url = {
      longUrl: form.longUrl,
      shortUrl: "abcdef"
    } 
    console.log(pair)
    setUrls([pair, ...urls]);
    setForm({ ...form, longUrl: "" })
  }

  const validateUrl = (url: string) => {
    return urlRegex.test(url);
  }

  return (
    <>
      <div className='body'>
        <Header/>
        <Form form={form} warning={showWarning} handleURLInputChange={urlInputChange} handleSubmit={genShortUrl} />
        <List urls={urls}/>
        <Footer/>
      </div>
    </>
  )
}

export default App
