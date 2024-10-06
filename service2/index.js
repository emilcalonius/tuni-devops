const express = require('express')
const ip = require('ip')
const { exec } = require('child_process')

const app = express()
const port = 3000

// Respond with container information
app.get('/', async function (_req, res) {
  const response = {
    ip: ip.address(),
    runningProcesses: await executeCommand('ps -ax'),
    diskSpace: await executeCommand('df'),
    timeSinceLastBoot: await getTimeSinceLastBoot(),
  }
  res.status(200).json(response)
})

/**
 * Spawn a shell and execute a command. Return the output.
 *
 * @param {string} command the command to execute
 * @returns {Promise<string>} a promise that resolves to the output of the command
 */
async function executeCommand(command) {
  return new Promise((resolve) => {
    exec(command, (error, stdout) => {
      if (error) {
        // Node couldn't execute the command
        console.error(`Error executing the '${command}' command:`, error)
        return
      }

      resolve(stdout)
    })
  })
}

/**
 * Get time since last boot by checking first process start time
 *
 * @returns {Promise<string>} string representing the hours, minutes, and seconds since last reboot
 */
async function getTimeSinceLastBoot() {
  const out = await executeCommand('stat /proc/1')
  const bootTime = new Date(out.split('Change: ')[1].split('.')[0])
  const currentTime = new Date()
  const hours = Math.floor(Math.abs(bootTime - currentTime) / (60 * 60 * 1000))
  const minutes = Math.floor(
    Math.abs(bootTime - currentTime) / (60 * 1000) - hours * 60
  )
  const seconds = Math.floor(
    Math.abs(bootTime - currentTime) / 1000 - hours * 60 * 60 - minutes * 60
  )
  return hours + ' hours, ' + minutes + ' minutes, ' + seconds + ' seconds'
}

app.listen(port, () => {
  console.log(`Listening on port ${port}`)
})
