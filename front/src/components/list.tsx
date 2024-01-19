import { Url } from '../interfaces/url';

interface Props {
  urls: Url[]
}

const List: React.FC<Props> = ({ urls }) => {
  return (
    <div className="container">
      <div className="list">
        {urls.map((item, idx) => (
          <div className="item" key={idx}>
              <a onClick={() => window.open(item.longUrl)} className="long-url">
                {item.longUrl}
              </a>
              <div className="short-url-container">
                <div className="tooltip">
                  <span>{item.shortUrl}</span>
                  <button className="copy" onClick={() => navigator.clipboard.writeText(item.shortUrl)}>Copy</button>
                </div>
              </div>
          </div>
        ))}
      </div>
    </div>
  )
}

export default List;
