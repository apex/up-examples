import os
import sys

from django.conf import settings
from django.conf.urls import url
from django.http import HttpResponse
from django.core.management import execute_from_command_line

# setup settings
DEBUG = os.environ.get('DEBUG', True)
SECRET_KEY = os.environ.get(
    'SECRET_KEY', 'A-random-and-secure-secret-key.\nKeep-this-very-SAFE!')
ALLOWED_HOSTS = os.environ.get('ALLOWED_HOSTS', '*').split(',')
settings.configure(
    DEBUG=DEBUG,
    SECRET_KEY=SECRET_KEY,
    ALLOWED_HOSTS=ALLOWED_HOSTS,
    ROOT_URLCONF=__name__,
    # Your other preferred django settings can be added here.
)


# setup urls/routing and views
def home_view(request):
    return HttpResponse('<h1>Hello from django and up!</h1>')


urlpatterns = (
    url(r'^$', home_view),
    # Your other preferred django urls can be added here.
)

if __name__ == "__main__":
    execute_from_command_line(sys.argv)
