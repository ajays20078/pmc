#cloud-config 
fqdn: __FQDN__
hostname: __HOSTNAME__
users: 
  - 
    name: krishnanvr
    shell: /bin/bash
    ssh_authorized_keys: 
      - ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDWSxKnWtFUW18woqZ/DTLb0Tv9GEI6r6zYuWZtHLshVLHux1wPwMmYUanuBkBzQODUdncoLaVeanp9SnRkeB04UgVsF1LJON9ca/3pRiXvNImUBN5p2DRSgKBOztocOjFKnHcK1q4oEQf/v6Pnc1bvtJaK0qj+vfNxt1h6xWKlcBTcXA1ZUNbq5oZel9pa311YxxWOL+kB7gpzNdX3MhAWvZAEsBK48sYq1hGLten6gMObgwQjrytS6tfJGd6F7Ybj1LGIKizOy9d2Ne61oMxI1gj8Cq07jRh3P4tSljnCXrqgHsjbvFbTKCd44fb1//cunqVplGfS3jl5TzCj7ti7 krishnan@phonepe
    sudo: 
      - ALL=(ALL) NOPASSWD:ALL
  - 
    name: nandudheeraj
    shell: /bin/bash
    ssh_authorized_keys: 
      - ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDVKyFJXe/wZIX4QJhsbgyrzLvOrMrGBUSH1J0eCDXAqfr6A6yuK6Ers9jB5Vstexr+H1p26itWRvpMyepqRV/Cg+mOOhRsaDi6CRgrEF+Sv0I4LC8QFc9lxI8RdCefzCbATCJQeRPZAva6D+loK7BIrNaul11nB/3HI0U3sUdgMFjifoFHnbaje6kT9bJwc10YLoIO1mLoX3vrHubmqZ3cjkudeotaUiUvGCFwBodzssZykNje1OqAfFEnpvmY8nHyWe6IgOVcuQuZWaAHC32Bysh5swI2jTrtAyrdqjBfNmd93Chwz6lsEhy2v76Li67EFbhOFoKJD8qcdRJpqYwX alapatidheeraj@Nandhus-MacBook-Pro.local
    sudo: 
      - ALL=(ALL) NOPASSWD:ALL
  - 
    name: bengineer
    shell: /bin/bash
    ssh_authorized_keys: 
      - ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDaMhETmG98xxvVoXQ52qHDcjLUcaawzdLNxMIuUe6F41uUZW3bvvl6j+qN7UbxP6pNMV+8IjQI3huZzrfSPhxt1uYVjIuxta1S6YbveactkkPNThmFOZxbcuoU/mPyR7i2zs3hdAjruVT+gHK6Pp/3KcUQZ82/YqKi4JI7iVPzuUtN4oppYa4qTfsoRXJ4Z6OxcQIguynvLrw2GAHV+t74TqLBlY94HjDqhHiph5ZpmL9quuFJ1qCGv33Bw+6t2tcalr7JOoCsN7R2F9lm8W8WAx1rrrFhwGtIml1GC2bQ9GMd/DZ8VdubrtleNZ0xplslGYo8fBW0TTrEu2skzc2f burzine@gmail.com
    sudo: 
      - ALL=(ALL) NOPASSWD:ALL