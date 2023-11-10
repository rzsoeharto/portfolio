import PropTypes from "prop-types";
import { stringFormat } from "../utils/formatter";

function SectionSelection({ index, sectionSelection, setImageArray }) {
  function handleChange(e) {
    sectionSelection.Content = e.target.textContent;
  }

  async function handleFileChange(event) {
    const file = event.target.files[0];

    const fileNameFormatted = stringFormat(file.name);
    let path = `https://firebasestorage.googleapis.com/v0/b/portfolio-project-6ac0e.appspot.com/o/images%2F${fileNameFormatted}?alt=media`;

    setImageArray((prevState) => [...prevState, file]);
    sectionSelection.Content = path;
    return;
  }

  let sectionType;

  switch (sectionSelection.SectionType) {
    case "Paragraph":
      sectionType = (
        <>
          <label>Paragraph</label>
          <span
            key={index}
            type="text"
            role="textbox"
            className="ParagraphBlock block w-full bg-white p-5 min-w-[740px] max-w-[740px] resize-y rounded focus:outline-none"
            contentEditable={true}
            onInput={handleChange}
          />
          <p className="min-w-[740px] max-w-[740px]">
            {sectionSelection.Content}
          </p>
        </>
      );
      break;

    case "Image":
      sectionType = (
        <>
          <img src={sectionSelection.Content} alt="" />
          <label>Upload an image</label>
          <input
            className="ParagraphBlock block w-full bg-white p-5 min-w-[740px] max-w-[740px] resize-y rounded focus:outline-none"
            type="file"
            placeholder="Upload an image"
            onChange={handleFileChange}
          />
        </>
      );
      break;

    case "CodeBlock":
      sectionType = (
        <>
          <label>Code</label>
          <span
            key={index}
            type="text"
            role="textbox"
            className="CodeBlock block w-full bg-white p-5 min-w-[740px] max-w-[740px] resize-y rounded focus:outline-none"
            contentEditable
            onInput={handleChange}
          />
        </>
      );
      break;

    default:
      break;
  }
  return <div className="flex flex-col">{sectionType}</div>;
}

SectionSelection.propTypes = {
  index: PropTypes.number,
  sectionSelection: PropTypes.object,
  setImageArray: PropTypes.func,
};

export default SectionSelection;
