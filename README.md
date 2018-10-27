# IGC track viewer extended - Clock trigger
Clock trigger part of assignment 2 in the course IMT2681-2018 (Cloud Technologies) at NTNU Gjøvik.

This a standalone application that checks the [paragliding API](https://github.com/mats93/paragliding) for new tracks added to the system every 10 min.

If new tracks are added, a Slack webhook is notified about the changes.

***

## Deployment:
 * The application is deployed in an OpenStack IaaS platform hosted at NTNU Gjøvik.
 * The infrastructure is created with OpenStack HEAT.
 * The application is running in a Docker Container.
 * The Docker server is created with a boot-script and the container with a docker compose file.

***

## Additional information:

Created by Mats Ove Mandt Skjærstein, 2018
