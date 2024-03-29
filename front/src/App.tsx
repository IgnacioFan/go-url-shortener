import "./App.css";
import { useState } from "react";
import Header from "./components/header";
import Form from "./components/form";
import Footer from "./components/footer";
import List from "./components/list";
import { createShortUrl } from "./apis/shortUrl";
import { Url } from "./interfaces/url";
import { formData } from "./interfaces/form";

const urlRegex = /^(https|ftp):\/\/(-\.)?([^\s/?\.#-]+\.?)+([^\s]*)$/i;
// avoid recursively redirections
const specificDomainRegex = /^(https:\/\/)?(localhost|google\.com)\b/i;


function App() {
  const [form, setForm] = useState<formData>({
    longUrl: ""
  });

  const [showWarning, setShowWarning] = useState<string>("");

  const [urls, setUrls] = useState<Url[]>([])

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
    createShortUrl(form.longUrl).then((responseData) => {
      let pair: Url = {
        longUrl: form.longUrl,
        shortUrl: responseData.data
      } 
      console.log(pair)
      setUrls([pair, ...urls]);
      setForm({ ...form, longUrl: "" })
      setShowWarning("")
    }).catch((err) => {
      // track the error's details
      setShowWarning("Unexpected error:" + err.message)
    })
  }

  const validateUrl = (url: string) => {
    return urlRegex.test(url) && !specificDomainRegex.test(url);
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
