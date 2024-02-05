import { formData } from "../interfaces/form";

interface Props {
  form: formData,
  warning: string,
  handleURLInputChange: (event: React.ChangeEvent<HTMLInputElement>) => void
  handleSubmit: (event: React.FormEvent<HTMLFormElement>) => void
}

const Form: React.FC<Props> = ({ form, warning, handleURLInputChange, handleSubmit }) => {
  return (
    <div className="container">
      <form onSubmit={handleSubmit} className="form">
        <div className="form-container">
          <div className="input-label">
            Long URL:
          </div>
          <input type="text" name="longUrl" className="input-field" value={form.longUrl} onChange={handleURLInputChange} />
          <button className="submit" type="submit">Submit</button>
        </div>
        <div className="form-container">
          {warning !== "" && (
            <div className="warning">
              <div style={{color: "red"}}>{warning}</div>
            </div>
          )}
        </div>
      </form>
    </div>
  )
}

export default Form;
