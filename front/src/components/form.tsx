import { useState } from "react";

interface FormData {
  url: string;
}

const Form: React.FC = () => {
  const [formData, setFormData] = useState<FormData>({ url: "" });

  function handleURLInputChange(event: React.ChangeEvent<HTMLInputElement>) {
    const { name, value } = event.target;
    setFormData({ ...formData, [name]: value });
  }

  function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();
    console.log(formData);
  }

  return (
    <div className="container">
      <form onSubmit={handleSubmit} className="form">
        <div className="input-label">
          <h3>Long URL:</h3>
          <input type="text" name="url" className="input-field" value={formData.url} onChange={handleURLInputChange} />
        </div>
        <button className="submit" type="submit">Submit</button>
      </form>
    </div>
  )
}

export default Form;
