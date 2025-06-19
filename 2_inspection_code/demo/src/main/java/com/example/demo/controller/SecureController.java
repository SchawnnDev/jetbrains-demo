package com.example.demo.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.util.HtmlUtils;

import java.io.*;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.List;

@Controller
@RequestMapping("/secure")
public class SecureController {

    @Autowired
    private JdbcTemplate jdbcTemplate;

    /**
     * Secure alternative to XSS
     * Properly escapes user input before rendering
     */
    @GetMapping("/xss")
    @ResponseBody
    public String xssSecure(@RequestParam String userInput) {
        // Escaping user input to prevent XSS
        String escapedInput = HtmlUtils.htmlEscape(userInput);
        return "<html><body><h1>Hello, " + escapedInput + "!</h1></body></html>";
    }

    /**
     * Secure alternative to SQL Injection
     * Uses parameterized queries to prevent SQL injection
     */
    @GetMapping("/sql")
    @ResponseBody
    public String sqlInjectionSecure(@RequestParam String username) {
        // Using parameterized query to prevent SQL injection
        String result;
        try {
            result = jdbcTemplate.queryForObject(
                    "SELECT email FROM users WHERE username = ?",
                    new Object[]{username},
                    String.class);
        } catch (Exception e) {
            return "Error: " + e.getMessage();
        }
        return "Email for " + username + " is: " + result;
    }

    /**
     * Secure alternative to Command Injection
     * Validates input and uses ProcessBuilder with a list of arguments
     */
    @GetMapping("/command")
    @ResponseBody
    public String commandInjectionSecure(@RequestParam String fileName) {
        // Input validation
        if (fileName.contains(";") || fileName.contains("|") || fileName.contains("&")) {
            return "Invalid input.";
        }

        StringBuilder output = new StringBuilder();
        try {
            // Using ProcessBuilder with separate arguments
            ProcessBuilder processBuilder = new ProcessBuilder("find", "/tmp", "-name", fileName);
            Process process = processBuilder.start();

            BufferedReader reader = new BufferedReader(
                    new InputStreamReader(process.getInputStream()));

            String line;
            while ((line = reader.readLine()) != null) {
                output.append(line).append("<br>");
            }
        } catch (IOException e) {
            return "Error: " + e.getMessage();
        }
        return "Search results: <br>" + output.toString();
    }

    /**
     * Secure alternative to Path Traversal
     * Validates the path to prevent directory traversal
     */
    @GetMapping("/file")
    @ResponseBody
    public String pathTraversalSecure(@RequestParam String fileName) {
        // Input validation - prevent traversal using ..
        if (fileName.contains("..")) {
            return "Invalid file name";
        }

        // Using Path normalization to prevent path traversal
        Path basePath = Paths.get("/app/public").normalize().toAbsolutePath();
        Path filePath = basePath.resolve(fileName).normalize().toAbsolutePath();

        // Additional validation to ensure the resolved path is within the base directory
        if (!filePath.startsWith(basePath)) {
            return "Access denied: path traversal detected";
        }

        StringBuilder content = new StringBuilder();
        try {
            BufferedReader reader = new BufferedReader(new FileReader(filePath.toFile()));

            String line;
            while ((line = reader.readLine()) != null) {
                content.append(line).append("<br>");
            }
            reader.close();
        } catch (IOException e) {
            return "Error: " + e.getMessage();
        }
        return "File content: <br>" + content.toString();
    }
}
