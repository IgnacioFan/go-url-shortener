import { formData } from "../interfaces/form";

interface Props {
  form: formData
  handleURLInputChange: (event: React.ChangeEvent<HTMLInputElement>) => void
  handleSubmit: (event: React.FormEvent<HTMLFormElement>) => void
}

const Form: React.FC<Props> = ({ form, handleURLInputChange, handleSubmit }) => {
  return (
    <div className="container">
      <form onSubmit={handleSubmit} className="form">
        <div className="input-label">
          <h3>Long URL:</h3>
          <input type="text" name="longUrl" className="input-field" value={form.longUrl} onChange={handleURLInputChange} />
        </div>
        <button className="submit" type="submit">Submit</button>
      </form>
    </div>
  )
}

export default Form;
