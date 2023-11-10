import PropTypes from "prop-types";

function AddSection({ setSectionModal, setSectionsData }) {
  function handleClick(e) {
    setSectionsData((prevState) => [
      ...prevState,
      {
        SectionType: e.target.value,
      },
    ]);
    setSectionModal(false);
  }

  return (
    <>
      <div className="flex flex-row w-[684px] h-[60px] bg-[#d9d9d9] rounded text-xl place-content-center justify-around px-20">
        <button
          className="bg-transparent hover:text-[#FFA360]"
          value="Paragraph"
          onClick={handleClick}
        >
          Paragraph
        </button>
        <div className="border-l border-zinc-500 h-1/2 self-center" />
        <button
          className="bg-transparent hover:text-[#FFA360]"
          value="Image"
          onClick={handleClick}
        >
          Image
        </button>
        <div className="border-l border-zinc-500 h-1/2 self-center" />
        <button
          className="bg-transparent hover:text-[#FFA360]"
          value="CodeBlock"
          onClick={handleClick}
        >
          Code
        </button>
      </div>
      <button
        onClick={() => {
          setSectionModal(false);
        }}
        className="font-bold bg-transparent h-full w-14"
      >
        X
      </button>
    </>
  );
}

AddSection.propTypes = {
  setSectionModal: PropTypes.func,
  setSectionsData: PropTypes.func,
};

export default AddSection;
