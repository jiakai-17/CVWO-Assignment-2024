import sampleThreads from "@/models/thread/SampleThreads";
import ThreadComponent from "@/components/ThreadComponent";
import Typography from "@mui/material/Typography";

export default function Page({ params }: Readonly<{ params: { id: string } }>) {

  const threadToDisplay = sampleThreads.find(thread => thread.id === params.id);

  if (threadToDisplay === undefined) {
    return (
      <Typography variant="h3">Thread not found</Typography>
    );
  }

  return (
    <ThreadComponent {...threadToDisplay} />
  );
}
