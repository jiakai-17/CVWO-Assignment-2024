// Defines a Thread type for the application.
export type Thread = {
  id: string;
  title: string;
  body: string;
  creator: string;
  created_time: Date;
  updated_time: Date;
  num_comments: number;
  tags: string[];
}

export default Thread;
