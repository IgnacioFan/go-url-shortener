import { useState } from "react";

interface Url {
  longUrl: string;
  shortUrl: string;
}

const List: React.FC = () => {
  const [urlList, setUrlList] = useState<Url[]>([
    {
      longUrl: "https://www.udemy.com/",
      shortUrl: "abc",
    },
    {
      longUrl: "https://www.kkday.com/zh-tw",
      shortUrl: "edc",
    }
  ]);

  return (
    <div className="container">
      <div className="list">
        {urlList.map((item, idx) => (
          <div className="item" key={idx}>
              <a onClick={() => window.open(item.longUrl)} className="long-url">
                {item.longUrl}
              </a>
              <div className="short-url-container">
                <div className="tooltip">
                  <span>{item.shortUrl}</span>
                  <button className="copy" onClick={() => console.log("copied short url " + item.shortUrl)}>Copy</button>
                </div>
              </div>
          </div>
        ))}
      </div>
    </div>
  )
}

export default List;
