import PropTypes from "prop-types";

function PostSection({ sectionData }) {
  const condition = sectionData.SectionType;

  let sectionContent;

  switch (condition) {
    case "Paragraph":
      sectionContent = <p className="text-xl">{sectionData.Content}</p>;
      break;

    case "Image":
      sectionContent = (
        <img
          src={sectionData.Content}
          alt="Content"
          className="max-w-[1080px]"
        />
      );
      break;

    case "CodeBlock":
      sectionContent = (
        <code className="w-full bg-gray-600 text-white rounded-md p-1">
          {sectionData.Content}
        </code>
      );
      break;

    default:
      sectionContent = <div>Something went wrong</div>;
      break;
  }

  return <>{sectionContent}</>;
}

PostSection.propTypes = {
  sectionData: PropTypes.object,
};

export default PostSection;
