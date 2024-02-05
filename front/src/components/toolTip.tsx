import { useState } from "react";

interface Props {
  copyValue: string
}

const ToolTip: React.FC<Props> = ({copyValue}) => {
  const [showTooltip, setShowTooltip] = useState<boolean>(false);

  const handleCopyClick = (copyValue: string) => {
    setShowTooltip(true);
    navigator.clipboard.writeText(copyValue)
    setTimeout(() => setShowTooltip(false), 1000);
  }
  return (
    <>
      <button className="tooltip-button" onClick={() => handleCopyClick(copyValue)}>Copy</button>
      {
        showTooltip && (
          <div className="tooltip">
            Copied!
          </div>
        )
      }
    </>
  )
}

export default ToolTip;
