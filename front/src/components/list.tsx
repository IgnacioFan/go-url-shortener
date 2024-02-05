import { Url } from "../interfaces/url";
import ToolTip from "./toolTip";

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
                <div className="tooltip-container" >
                  <span>{item.shortUrl}</span>
                  <ToolTip copyValue={item.shortUrl}/>
                </div>
              </div>
          </div>
        ))}
      </div>
    </div>
  )
}

export default List;
