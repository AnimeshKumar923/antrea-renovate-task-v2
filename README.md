# Prerequisite task for Antrea (LFX mentorship, Term 2, 2025)

This repository is for demonstrating the task given for LFX mentorship for [Antrea project](https://github.com/antrea-io/antrea).

- Project: [Replace Dependabot with Renovate for automatic dependency updates](https://github.com/cncf/mentoring/tree/main/programs/lfx-mentorship/2025/02-Jun-Aug#replace-dependabot-with-renovate-for-automatic-dependency-updates)
- LFX URL: [LFX mentorship](https://mentorship.lfx.linuxfoundation.org/project/62d69fd5-6c90-4ba1-b260-a5dc247fc3cf)

## Description
- This project is designed to showcase how Renovate Bot can be configured to automatically detect vulnerabilities in Go module dependencies and raise PRs to update them. By intentionally including a known vulnerable version of a Go package, we observe how Renovate Bot identifies the issue and suggests a fix. It imports the vulnerable package but doesn't perform any significant operations. It's solely for demonstration purposes.
- I've kept the implementation simple to focus on the Renovate bot's functionality.

## How to run

1. Clone the repository:
  ```bash
  git clone https://github.com/AnimeshKumar923/antrea-renovate-task-v2/
  ```

2. Change directory to the cloned repository:
  ```bash
  cd antrea-renovate-task-v2
  ```
3. Run the Go program:
  ```bash
  go run main.go
  ```
4. Check the output:
  ```bash
  Demo using golang.org/x/crypto
  ```

![Screenshot from 2025-05-22 13-30-53](https://github.com/user-attachments/assets/9a9ed336-00ba-4207-855b-9315253afaec)


## Vulnerability
### Description
The SSH transport protocol with certain OpenSSH extensions, found in OpenSSH before 9.6 and other products, allows remote attackers to bypass integrity checks such that some packets are omitted (from the extension negotiation message), and a client and server may consequently end up with a connection for which some security features have been downgraded or disabled, aka a Terrapin attack. _(Source: CVE-2023-48795)_

### Impact 
By weakening the SSH session's security, sensitive data transmitted during the session could be exposed or modified.

### Vulnerability Metrics
- Vector:  CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:U/C:N/I:H/A:N
- Base Score:  5.9 (MEDIUM)
- Attack Vector:  Network
- Attack Complexity:  High
> [!TIP]
> Read more about the vulnerability [here](https://nvd.nist.gov/vuln/detail/cve-2023-48795).

## Renovate Bot Details
- Located at `.github/renovate.json`

### JSON Configuration
```json
{
  "$schema": "https://docs.renovatebot.com/renovate-schema.json",
  "extends": ["config:base"],
  "vulnerabilityAlerts": {
    "enabled": true
  },
  "assignAutomerge": true,
  "assigneesFromCodeOwners": true,
  "postUpdateOptions": ["gomodTidy"],
  "baseBranches": ["main"],
  "packageRules": [
    {
      "matchManagers": ["gomod"],
      "labels": ["area/security", "dependencies"],
      "automerge": false,
      "commitMessageTopic": "{{depName}}",
      "commitMessageExtra": "to {{newVersion}}"
    }
  ]
}
```
Config TLDR (kind of):
- Enables vulnerability alerts via GitHub Security Advisories.

- Prevents auto-merging of Go deps.

- Cleans go.mod after each update.

- Labels and structures commits cleanly for review.

- Works specifically and only on the main branch (more can be added by extending `baseBranches` or `matchBaseBranches` depending upon our use-cases).

- Pulls assignees from `CODEOWNERS` to keep responsibilities clear.

### Bot Workflow
  1. **Renovate Bot Scans Dependencies**: It checks for known vulnerabilities in your dependencies using sources like GitHub Security Advisories, etc.

  2. **PR Creation**: Upon detecting a vulnerability, Renovate Bot automatically creates a PR to update the affected dependency to a secure version.

  3. **Review and Merge**: You can review the PR, which includes details about the vulnerability, and merge it to apply the fix.

### Renovate Configuration (on website)
![renovate-website-settings](https://github.com/user-attachments/assets/23ff4e52-cdb3-4fa6-a3bf-741045c920f2)


### Renovate Bot PR
- The PR created by Renovate Bot will include details about the vulnerability and affected dependency. It will also include a link to the relevant CVE report for further information.
- The PR will be labeled with `area/security` and `dependencies` to categorize it appropriately.
- The PR will not be automatically merged, allowing for manual review and testing before applying the changes.
- The commit message will follow the format `{{depName}} to {{newVersion}}`, providing a clear indication of the changes made.
- The PR will be assigned to the code owners specified in the repository, ensuring that the right people are notified about the changes.

## Future Improvements

In the development of this feature some ideas which can help are:

- A schedule to check for the vulnerabilities in the dependencies and update them.
- Maybe a notification system to alert the maintainers when a vulnerability is found (in the Slack channel, via GH notifications or e-mail as deemed fit).
- [ignorePaths](https://docs.renovatebot.com/configuration-options/#ignorepaths) can be used to ignore the folders and files that are not relevant for vulnerability scanning.
- Usage of semantic commit for consistent and clear commit messages.
- Define more active branches _(like `release-2.2`, `release-2.1`, as mentioned in the [original issue](https://github.com/antrea-io/antrea/issues/6934))_ on which we can scan for vulnerabilities.

> [!TIP]
> More practices can be adopted by analyzing projects that use Renovate bot and have a good security policy. (like cilium)

## References
- [VulnerabilityAlerts Docs](https://docs.renovatebot.com/configuration-options/#vulnerabilityalerts)
- [CVE Report](https://nvd.nist.gov/vuln/detail/cve-2023-48795)
- [Cilium Renovate Bot Config](https://github.com/cilium/cilium/blob/main/.github/renovate.json5)
- [Original Detailed Issue](https://github.com/antrea-io/antrea/issues/6934)
