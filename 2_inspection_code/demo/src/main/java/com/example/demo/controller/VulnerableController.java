package com.example.demo.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.*;

import java.io.*;
import java.sql.Connection;
import java.sql.Statement;

@Controller
@RequestMapping("/vulnerable")
public class VulnerableController {

    @Autowired
    private JdbcTemplate jdbcTemplate;

    /**
     * Exemple de vulnérabilité XSS
     * L'entrée utilisateur est directement rendue dans la vue sans échappement approprié
     */
    @GetMapping("/xss")
    @ResponseBody
    public String xssVulnerable(@RequestParam String userInput) {
        // Vulnérable au XSS : insertion directe de l'entrée utilisateur dans la réponse
        return "<html><body><h1>Hello, " + userInput + "!</h1></body></html>";
    }

    /**
     * Exemple de vulnérabilité d'injection SQL
     * L'entrée utilisateur est directement concaténée dans la requête SQL
     */
    @GetMapping("/sql")
    @ResponseBody
    public String sqlInjectionVulnerable(@RequestParam String username) {
        // Vulnérable à l'injection SQL : insertion directe de l'entrée utilisateur dans la requête SQL
        String result = "";
        try {
            result = jdbcTemplate.queryForObject(
                    "SELECT email FROM users WHERE username = '" + username + "'",
                    String.class);
        } catch (Exception e) {
            return "Error: " + e.getMessage();
        }
        return "Email for " + username + " is: " + result;
    }

    /**
     * Exemple de vulnérabilité d'injection de commande
     * L'entrée utilisateur est directement utilisée dans une commande système
     */
    @GetMapping("/command")
    @ResponseBody
    public String commandInjectionVulnerable(@RequestParam String fileName) {
        StringBuilder output = new StringBuilder();
        try {
            // Vulnérable à l'injection de commande : utilisation directe de l'entrée utilisateur dans une commande
            Process process = Runtime.getRuntime().exec("find /tmp -name " + fileName);
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
     * Exemple de vulnérabilité de traversée de chemin
     * L'entrée utilisateur est directement utilisée dans le chemin du fichier
     */
    @GetMapping("/file")
    @ResponseBody
    public String pathTraversalVulnerable(@RequestParam String fileName) {
        StringBuilder content = new StringBuilder();
        try {
            // Vulnérable à la traversée de chemin : utilisation directe de l'entrée utilisateur dans le chemin du fichier
            File file = new File("/app/public/" + fileName);
            BufferedReader reader = new BufferedReader(new FileReader(file));

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
