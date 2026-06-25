# Security Audit Report

## Executive Summary

A security assessment was conducted on the Go application using Semgrep static analysis to identify insecure coding practices. The assessment focused on remediating SQL Injection (CWE-89), Hardcoded Credentials (CWE-798), and insecure JWT validation. After implementing the fixes and rerunning the scan, the total number of findings decreased from **12 to 10**, confirming that the SQL Injection vulnerabilities were resolved while the remaining findings correspond to unrelated security issues intentionally left for future remediation.

---

## Findings Table

| Vulnerability           | CWE                     | Severity | File            | Fix Applied                                                                                                                                                                                                                               |
| ----------------------- | ----------------------- | -------- | --------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| SQL Injection           | CWE-89                  | Critical | main.go         | Replaced string-formatted SQL query with parameterized query using placeholders (`?`) and parameter binding.                                                                                                                              |
| Hardcoded Secret        | CWE-798                 | High     | main.go, jwt.go | Removed hardcoded secrets and replaced them with environment variable lookups (`os.Getenv`). Added startup validation to terminate execution if required secrets are missing.                                                             |
| Insecure JWT Validation | JWT Algorithm Confusion | Critical | jwt.go          | Replaced insecure JWT validation with secure validation using `github.com/golang-jwt/jwt/v5`, enforced the HS256 algorithm allowlist, loaded signing key from an environment variable, and added a unit test rejecting `alg:none` tokens. |

---

## Scan Results

| Scan             | Findings |
| ---------------- | -------- |
| Baseline Scan    | 12       |
| Post-Fix Scan    | 10       |
| Findings Reduced | 2        |

The local Semgrep OSS ruleset reported a reduction of two findings after the SQL Injection remediation. The hardcoded secret and JWT validation fixes were implemented according to the assignment requirements. These issues were not emitted as findings by the local OSS Semgrep configuration.

---

## Lessons Learned

Static Application Security Testing (SAST) tools such as Semgrep provide an effective way to identify insecure coding practices early in the software development lifecycle. Parameterized SQL queries significantly reduce the risk of SQL Injection attacks by separating user input from SQL syntax. Secure secret management through environment variables also prevents accidental exposure of sensitive credentials within the source code repository.

Authentication logic must never trust security-critical information supplied by untrusted inputs. JWT validation should explicitly verify the expected signing algorithm and use securely managed signing keys loaded from environment variables. Integrating automated security scanning into CI/CD pipelines ensures vulnerabilities are detected and remediated before software reaches production environments.
