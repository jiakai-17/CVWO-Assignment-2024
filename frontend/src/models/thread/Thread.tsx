// Defines a Thread type for the application.
export type Thread = {
  id: string;
  title: string;
  body: string;
  creator: string;
  created_time: string;
  updated_time: string;
  num_comments: number;
  tags: string[];
};

export default Thread;
