import SearchIcon from "@mui/icons-material/Search";
import Button from "@mui/material/Button";

export default function SearchButton(
  props: Readonly<{
    size: "small" | "large";
    onClick?: () => void;
  }>,
) {
  return (
    <Button
      variant="contained"
      disableElevation
      size={props.size}
      endIcon={props.size === "large" ? <SearchIcon /> : undefined}
      onClick={props.onClick}
    >
      {props.size === "large" ? "Search" : <SearchIcon />}
    </Button>
  );
}
