import { createLogger, format, transports } from "winston";

// Create a structured logger using Winston
const logger = createLogger({
  level: process.env.LOG_LEVEL!,
  format: format.combine(
    format.timestamp(), // Add a timestamp
    format.json() // Use JSON format for structured logs
  ),
  transports: [
    new transports.Console(), // Log to the console
  ],
});

// Export the logger
export default logger;
