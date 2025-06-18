#ifndef OBSERVABLE_H
#define OBSERVABLE_H

#include"./../Observer/Observer.hpp"

class Observable{
    
    public:

    virtual void addObserver(Observer* observer)=0;
    virtual void removeObserver(Observer * observer)=0;

    virtual void notifyObservers()=0;

};


#endif