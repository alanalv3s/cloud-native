import express, { Request, Response } from "express";
import logger from "./util/logger";
// eslint-disable-next-line @typescript-eslint/no-require-imports
require("dotenv").config();

const app = express();
const port = process.env.PORT || 3000;

app.get("/health", (req: Request, res: Response) => {
  res.status(200).json({
    status: "OK",
    message: "Service is running",
  });
});

app.listen(port, () => {
  logger.info(`Server is running on port ${port}`);
});

// Middleware to log requests
app.use((req: Request, res: Response, next) => {
  logger.info(`Incoming request: ${req.method} ${req.url}`); // Log request details
  next();
});

// Add an error handler
app.use((err: Error, req: Request, res: Response, next: Function) => {
  logger.error(`Error occurred: ${err.message}`); // Log error message
  res.status(500).send("Internal Server Error");
});
