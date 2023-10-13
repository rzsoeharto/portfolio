import PropTypes from "prop-types";
import { postStorage } from "../stores/useStore";

function SectionSelection({ index, sectionSelection }) {
  const { updateSectionContent } = postStorage();

  function handleChange(e) {
    sectionSelection.Content = e.target.textContent;
    updateSectionContent(index, sectionSelection.Content);
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
          <label>Image</label>
          <input
            className="ParagraphBlock block w-full bg-white p-5 min-w-[740px] max-w-[740px] resize-y rounded focus:outline-none"
            type="file"
            placeholder="Upload an image"
          />
        </>
      );
      break;

    case "CodeBlock":
      sectionType = (
        <>
          <label>Code</label>
          <span
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
};

export default SectionSelection;
